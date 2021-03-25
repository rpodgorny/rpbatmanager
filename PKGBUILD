# Maintainer: Radek Podgorny <radek@podgorny.cz>
pkgname=rpbatmanager-git
provides=('rpbatmanager')
conflicts=('rpbatmanager')
pkgver=0.0
pkgrel=1
pkgdesc="Radek Podgorny's battery manager"
arch=('any')
url="https://github.com/rpodgorny/rpbatmanager"
#license=('PSF')
depends=('babashka')
makedepends=('git')
#options=(!emptydirs)
#backup=('etc/faddnsc.conf')
source=("$pkgname::git+https://github.com/rpodgorny/rpbatmanager")
md5sums=('SKIP')

pkgver() {
	cd "$srcdir/$pkgname"
	printf "r%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short HEAD)"
}

package() {
	cd "$srcdir/$pkgname"
	#python setup.py install --root="$pkgdir"/ --optimize=1

	install -m 0644 rpbatmanager $pkgdir/usr/bin
	#install -d $pkgdir/usr/lib/systemd/system
	install -m 0644 rpbatmanager.service $pkgdir/usr/lib/systemd/system/

	#install -d $pkgdir/etc
	#install -m 0644 etc/faddnsc.conf $pkgdir/etc/
}
