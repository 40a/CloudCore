CloudCore Debian Package
===========================

To create an CloudCore DEB package:

    git clone git://github.com/CloudPac/CloudCore.git
    cd CloudCore
    make deb

The debian package file will be placed in the `../` directory. This can then be added to an APT repository or installed with `dpkg -i <package-file>`.

Note that `dpkg -i` does not resolve dependencies.

To install the CloudCore DEB package and resolve dependencies:

    sudo dpkg -i <package-file>
    sudo apt-get -fy install

Or, if you are running Debian Stretch (or later) or Ubuntu Xenial (or later):

    sudo apt install <package-file>