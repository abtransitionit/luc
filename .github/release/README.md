# Todo
Should adopt a per release file
```
/
├── releases/
│   ├── v0.1.0.md
│   ├── v0.2.0.md
│   └── template.md
├── ...
```

Then in release.yaml:
```yaml
body_path: releases/${{ env.TAG_NAME }}.md
```