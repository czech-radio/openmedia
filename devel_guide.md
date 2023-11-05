# DEVELOPMENT GUIDE AND DISCUSSION
## WHICH GOLANG VERSION TO LOCK?
*2023-11-05*
-   arch_actual: go version go1.21.3 linux/amd64
-   alpine_3.18 (./deploy/dockerfile_devel.yml): go version go1.20.10 linux/amd64

## LOGGING
https://betterstack.com/community/guides/logging/best-golang-logging-libraries/
-   slog: new bulit-in logging in Go 1.21
-   zerolog: fastest
-   zap: fast yet flexible
