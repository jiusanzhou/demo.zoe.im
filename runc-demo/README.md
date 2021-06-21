

```bash
docker export alpine > rootfs.tar
mkdir rootfs
tar -xf rootfs.tar -C ./rootfs
rm rootfs.tar
```