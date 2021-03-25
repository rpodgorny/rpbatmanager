# Maintainer: Radek Podgorny <radek@podgorny.cz>
pkgname=rpbatmanager-git
provides=('rpbatmanager')
conflicts=('rpbatmanager')
pkgver=r2.7a788fd
pkgrel=1
pkgdesc="Radek Podgorny's battery manager"
arch=('any')
url="https://github.com/rpodgorny/rpbatmanager"
#license=('PSF')
depends=('babashka')
makedepends=('git')
#options=(!emptydirs)
#backup=('etc/rpbatmanager.conf')
source=("$pkgname::git+https://github.com/rpodgorny/rpbatmanager")
md5sums=('SKIP')

pkgver() {
	cd "$srcdir/$pkgname"
	printf "r%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short HEAD)"
}

package() {
	cd "$srcdir/$pkgname"

	install -D -t $pkgdir/usr/bin rpbatmanager
	install -D -m 0644 -t $pkgdir/usr/lib/systemd/system/ rpbatmanager.service
}
