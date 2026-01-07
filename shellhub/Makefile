PORTNAME=	shellhub
DISTVERSIONPREFIX=	v
DISTVERSION=	0.21.5
DISTVERSIONSUFFIX=
CATEGORIES=	sysutils
MASTER_SITES=	https://github.com/shellhub-io/shellhub/releases/download/${DISTVERSIONPREFIX}${DISTVERSION}${DISTVERSIONSUFFIX}/
DISTNAME=	shellhub-agent
DIST_SUBDIR=	${DISTNAME}${DISTVERSION}${DISTVERSIONSUFFIX}

MAINTAINER=	henry.barreto@ossystems.com.br
COMMENT=	ShellHub Agent component
WWW=		https://shellhub.io/

LICENSE=	APACHE20

USES=		go
USE_RC_SUBR=	${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-X main.AgentVersion=${DISTVERSIONPREFIX}${DISTVERSION}${DISTVERSIONSUFFIX}"

WRKSRC=		${WRKDIR}/agent

.include <bsd.port.mk>
