%define name cloudcore
%define cloud_core_version $VERSION

Name:      %{name}
Version:   %{cloud_core_version}
Release:   1%{?dist}
Url:       http://github.com/CloudPac/CloudCore
Summary:   Comprehensive web-based automation framework and Centralized infrastructure management platform
License:   GPLv3
Group:     Tools/CloudCore
Source0:   cloudcore-%{version}.tar.gz
BuildRoot: %{_tmppath}/%{name}-%{version}-%{release}-buildroot

BuildArch: x86_64

# RHEL > 5
%if 0%{?rhel} && 0%{?rhel} > 5
Requires: ansible
Requires: git
Requires: subversion
Requires: mercurial
%endif

# FEDORA > 17
%if 0%{?fedora} >= 18
Requires: ansible
Requires: git
Requires: subversion
Requires: mercurial
%endif

# SuSE/openSuSE
%if 0%{?suse_version} 
Requires: ansible
Requires: git
Requires: subversion
Requires: mercurial
%endif

%description

Comprehensive web-based automation framework and Centralized infrastructure management platform,
provides role-based access control, job scheduling, inventory management.

%prep
%setup -q

%install

mkdir -p %{buildroot}%{_bindir}
mkdir -p %{buildroot}%{_sharedstatedir}/cloudcore/
mkdir -p %{buildroot}%{_mandir}/man1/
mkdir -p %{buildroot}/lib/systemd/system/
mkdir -p %{buildroot}%{_sysconfdir}/
cp -v bin/* %{buildroot}%{_bindir}
cp -v docs/man/*.1 %{buildroot}%{_mandir}/man1/
cp -v etc/cloud-core.conf %{buildroot}%{_sysconfdir}/
cp -rv lib/* %{buildroot}%{_sharedstatedir}/cloudcore/
cp -v systemd/cloudcored.service %{buildroot}/lib/systemd/system/

%clean
rm -rf %{buildroot}

%files
%defattr(-,root,root)
%{_bindir}/cloudcore*
%config(noreplace) %{_sysconfdir}/cloud-core.conf
%doc %{_mandir}/man1/cloudcore*
%{_sharedstatedir}/cloudcore/
/lib/systemd/system/cloudcored.service

%post
if ! getent group cloudcore > /dev/null; then
    useradd -rU cloudcore
fi

%changelog

* Fri Nov 25 2016 Gamunu Balagalla <gamunu.balagalla@outlook.com> - 0.1.5
- Release 0.1.5

* Fri Nov 25 2016 Gamunu Balagalla <gamunu.balagalla@outlook.com> - 0.1.4
- Release 0.1.4

* Mon Nov 21 2016 Gamunu Balagalla <gamunu.balagalla@outlook.com> - 0.1.3
- Release 0.1.3

* Thu Nov 10 2016 Gamunu Balagalla <gamunu.balagalla@outlook.com> - 0.0.12
- Release 0.0.12

* Tue Nov 8 2016 Gamunu Balagalla <gamunu.balagalla@outlook.com> - 0.1.0
- Release 0.1.0

* Thu Nov 3 2016 Gamunu Balagalla <gamunu.balagalla@outlook.com> - 0.0.11
- Release 0.0.11

* Thu Nov 3 2016 Gamunu Balagalla <gamunu.balagalla@outlook.com> - 0.0.10
- Release 0.0.10

* Wed Nov 2 2016 Gamunu Balagalla <gamunu.balagalla@outlook.com> - 0.0.9
- Release 0.0.9

* Tue Nov 1 2016 Gamunu Balagalla <gamunu.balagalla@outlook.com> - 0.0.8
- Release 0.0.8

* Mon Oct 31 2016 Gamunu Balagalla <gamunu.balagalla@outlook.com> - 0.0.7
- Release 0.0.7