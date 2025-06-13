# Process
- install kind, nerdctl, nerdctl, rootlesskit, slirp4netns
- export KIND_EXPERIMENTAL_PROVIDER=nerdctl
- export path to kind, nerdctl, nerdctl
  - sanitize PATH by eliminating doublons
- sudo apt-get install -y uidmap
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
	

Absolutely! Here’s a clear step-by-step doc on how to fix the rootless containerd problem caused by AppArmor restrictions, using standard file creation and commands:

# Explanation

* **AppArmor** is a Linux security module that restricts program capabilities.
* By default, it may block **rootlesskit** from creating user namespaces needed for rootless containers.
* The override file grants **rootlesskit** the `userns` capability explicitly.
* Restarting AppArmor reloads these changes, allowing rootlesskit to function properly.



# Reference
- https://kind.sigs.k8s.io/docs/user/quick-start/
- https://kind.sigs.k8s.io/docs/design/initial
- https://rootlesscontaine.rs/getting-started/common/
- https://rootlesscontaine.rs/
- https://github.com/mgoltzsche/slirp-cni-plugin?tab=readme-ov-file
- https://man.archlinux.org/man/extra/slirp4netns/slirp4netns.1.en