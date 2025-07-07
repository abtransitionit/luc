# Process
- install CLI(s):
  - **gocli**: kind, nerdctl, containerd, rootlesskit, slirp4netns
  - **dnfapt**: uidmap
- configure os services for
  - create file content in memory
  - save memory content to file
- export linux env variables:
  - **KIND_EXPERIMENTAL_PROVIDER**
    - KIND_EXPERIMENTAL_PROVIDER=nerdctl
  - **PATH** : will all gocli binaries
- sanitize PATH by eliminating doublons

-  Create an AppArmor Profile Override for rootlesskit
```bash
cat <<EOT | sudo tee "/etc/apparmor.d/usr.local.bin.rootlesskit.rootlesskit"
# ref: https://ubuntu.com/blog/ubuntu-23-10-restricted-unprivileged-user-namespaces
abi <abi/4.0>,
include <tunables/global>

/usr/local/bin/rootlesskit/rootlesskit flags=(unconfined) {
  userns,

  # Site-specific additions and overrides. See local/README for details.
  include if exists <local/usr.local.bin.rootlesskit.rootlesskit>
}
EOT
```
- sudo systemctl restart apparmor.service

```bash
sudo tee /etc/apparmor.d/usr.local.bin.rootlesskit.rootlesskit > /dev/null <<EOF
# Allow rootlesskit to create user namespaces (userns)
# Ref: https://ubuntu.com/blog/ubuntu-23-10-restricted-unprivileged-user-namespaces
abi <abi/4.0>,
include <tunables/global>

/usr/local/bin/rootlesskit/rootlesskit flags=(unconfined) {
  userns,

  # Site-specific additions and overrides. See local/README for details.
  include if exists <local/usr.local.bin.rootlesskit.rootlesskit>
}
EOF

```
- play 'containerd-rootless-setuptool.sh install'
	- check it created file '/run/user/1000/containerd-rootless'
- play 'nerdctl info'
- rootless kit
  - $ echo $XDG_RUNTIME_DIR (should output /run/user/1000)
- enable linger so containerd starts on boot for your user:
```bash
sudo loginctl enable-linger $(whoami)
```
	



# Toknow
**AppArmor** and **Selinux** allows both to improve linux Os security

Both :  
✔ **Restrict what programs can do** (beyond traditional file permissions).  
✔ **Enforce fine-grained access control** (e.g., which files a process can read/write).  
✔ **Prevent privilege escalation** and limit damage from compromised apps.  



| Feature          | **AppArmor** | **SELinux** |
|-----------------|------------|------------|
| **Developed By** | Canonical (Ubuntu) | NSA (later Red Hat) |
| **OS** | Ubuntu/Debian | RHEL/Fedora |
| **Configuration** | Path-based rules (`/usr/bin/foo`) | Label-based (`system_u:object_r:httpd_exec_t`) |
| **Ease of Use** | Simpler (uses plain-text profiles) | More complex (uses security contexts) |
| **Default in**  | Ubuntu, Debian, openSUSE | RHEL, Fedora, CentOS, Android |
| **Performance** | Lighter weight | More overhead (but more powerful) |
| **Policy Type** | **Name-based** (e.g., `/var/log/** rw`) | **Type Enforcement (TE) + Role-Based Access Control (RBAC)** |

## AppArmor
- A Linux security module that restricts program capabilities.
- By default, it may block **rootlesskit** from creating user namespaces needed for rootless containers.
- Uses **path-based rules** (e.g., allow `/usr/bin/nginx` to read `/var/log/nginx/*`).  
- Profiles are stored in `/etc/apparmor.d/`.  
- Example profile:  
  ```apparmor
  /usr/bin/firefox {
    /etc/firefox/* r,
    /home/*/.mozilla/** rw,
  }
  ```

- What it does
  * The override file grants **rootlesskit** the `userns` capability explicitly.

- Restarting AppArmor reloads these changes, allowing rootlesskit to function properly.


```bash
sudo aa-status          # Check status
sudo apparmor_parser -r /etc/apparmor.d/profile  # Reload a profile
sudo aa-enforce /usr/bin/firefox  # Enforce a profile
```


## SELinux
- Uses **security labels** (e.g., `httpd_t` can only access `httpd_log_t`).  
- Managed via `semanage`, `chcon`, `restorecon`.  
- Example rule:  
  ```bash
  chcon -t httpd_sys_content_t /var/www/html/index.html
  ```

```bash
sestatus                # Check status
ls -Z /var/www/html     # View security labels
restorecon -Rv /var/www # Fix labels
setenforce 0            # Temporarily disable (Permissive mode)
```

# User service Vs. System service

| Feature| User Service | System Service |
| -| - | - |
| Runs as         | Regular user                     | Root (or system-level)                           |
| Managed by      | `systemctl --user`               | `sudo systemctl`                                 |
| Config location | `~/.config/systemd/user/`        | `/etc/systemd/system/` or `/lib/systemd/system/` |
| Starts at boot? | Only if **lingering** is enabled | Yes, by default                                  |
| Scope           | Only affects the user            | Affects the whole system                         |


# Lingering
- allows to enable user services at startup (equivalent of system enable for system services)

```bash
# show infos for all user
loginctl show-user 

# show infos for user ubuntu
loginctl show-user ubuntu --property=Display
loginctl show-user ubuntu

# enable lingering
sudo loginctl enable-linger ubuntu

# disable lingering
sudo loginctl disable-linger ubuntu
```
ls /var/lib/systemd/linger
```bash
loginctl show-user ubuntu
```


# Reference
- https://kind.sigs.k8s.io/docs/user/quick-start/
- https://kind.sigs.k8s.io/docs/design/initial
- https://rootlesscontaine.rs/getting-started/common/
- https://rootlesscontaine.rs/
- https://github.com/mgoltzsche/slirp-cni-plugin?tab=readme-ov-file
- https://man.archlinux.org/man/extra/slirp4netns/slirp4netns.1.en