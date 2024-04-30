PORTNAME=	shellhub-agent
DISTVERSIONPREFIX=	v
DISTVERSION=	0.15.1
DISTVERSIONSUFFIX=	-rc.1
DISTNAME=	shellhub-agent
DIST_SUBDIR=	${DISTNAME}${DISTVERSION}${DISTVERSIONSUFFIX}
CATEGORIES=	sysutils
MASTER_SITES=	https://github.com/shellhub-io/shellhub/releases/download/${DISTVERSIONPREFIX}${DISTVERSION}${DISTVERSIONSUFFIX}/

MAINTAINER=	henry.barreto@ossystems.com.br
COMMENT=	ShellHub Agent component
WWW=		https://shellhub.io/

LICENSE=	APACHE20

USES=		go
GO_BUILDFLAGS+=	-ldflags="-X main.AgentVersion=${DISTVERSIONPREFIX}${DISTVERSION}${DISTVERSIONSUFFIX}"

WRKSRC=		${WRKDIR}/agent

.include <bsd.port.mk>
