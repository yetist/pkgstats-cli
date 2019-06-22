# Maintainer: Pierre Schmitz <pierre@archlinux.de>

pkgname=pkgstats
pkgver=2.4
pkgrel=1
pkgdesc='Submit a list of installed packages to the Arch Linux project'
url='https://pkgstats.archlinux.de/'
arch=('any')
license=('GPL')
depends=('bash' 'curl' 'pacman' 'sed' 'coreutils' 'systemd' 'grep')
source=(pkgstats.{sh,timer,service})
sha256sums=('4c96e0946c44435a3853fffd4fe99ae25aea3efa9a567b918eb9ca27c6cf9b52'
            '86207164a13640edb58657f16329f60f2d84d7d3e5b9336e48aa0d607906078e'
            '986608f2fff417693b663474db3f36b8fb2ae4eb111ad177c616ce02bb431b23')

package() {
	install -D pkgstats.sh "$pkgdir/usr/bin/pkgstats"
	install -Dt "$pkgdir/usr/lib/systemd/system" -m644 pkgstats.{timer,service}
	install -d "$pkgdir/usr/lib/systemd/system/timers.target.wants"
	ln -st "$pkgdir/usr/lib/systemd/system/timers.target.wants" ../pkgstats.timer
}
