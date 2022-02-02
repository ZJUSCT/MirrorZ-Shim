# MirrorZ-Shim
 Shim Layer for MirrorZ written in Go

## Configs
 Set Envs:
 - `MIRRORZ_SHIM_URL` origin status json url, default to https://mirrors.zju.edu.cn/api/mirrors
 - `MIRRORZ_SHIM_CACHE_TTL` cache TTL in minutes, default to 5

## CNAME File
`/app/configs/mirrorz-cname.json`

> Notes:
> 1. This json has been modified, remember to merge them when mirrorz has an updated one.
> 2. MirrorZ `cname.json`: https://github.com/mirrorz-org/mirrorz/blob/master/static/json/cname.json