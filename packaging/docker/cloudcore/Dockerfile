FROM ubuntu:16.04

# add our user and group first to make sure their IDs get assigned consistently, regardless of whatever dependencies get added
RUN groupadd -r cloudcore && useradd -r -m -g cloudcore cloudcore && \
    export DEBIAN_FRONTEND=noninteractive && \
    apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 6125E2A8C77F2818FB7BD15B93C4A3FD7BB9C367 && \
    apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 805A0463B9370AABA2C9D3192EE7B9E9BF6D1FEE && \
    apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 6125E2A8C77F2818FB7BD15B93C4A3FD7BB9C367 && \
    echo "deb http://ppa.launchpad.net/ansible/ansible/ubuntu xenial main" | tee -a /etc/apt/sources.list.d/ansible.list && \
    echo "deb http://ppa.launchpad.net/ansible/proot/ubuntu xenial main" | tee -a /etc/apt/sources.list.d/proot.list && \
    echo "deb http://ppa.launchpad.net/cloudpac/cloudcore/ubuntu xenial main" | tee -a /etc/apt/sources.list.d/cloud-core.list && \
    apt-get -q update && \
    apt-get -y -q install \
    ansible \
    git \
    subversion \
    mercurial \
    python-dev \
    libkrb5-dev \
    python-pip \
    curl \
    unzip \
    libssl-dev \
    libffi-dev \
    python-dev \
    proot \
    krb5-user && \
    # install python packages
    pip install -U \
    pip \
    wheel \
    setuptools && \
    pip install --upgrade \
    "pywinrm>=0.2.2" \
    "azure==2.0.0rc5" \
    pyrax \
    apache-libcloud \
    boto \
    kerberos \
    requests-kerberos \
    requests-credssp && \
    apt-get -q -y upgrade && apt-get -y -q --no-install-recommends install -f cloudcore && \
    #install terraform
    cd /tmp/ && curl -O https://releases.hashicorp.com/terraform/0.9.5/terraform_0.9.5_linux_amd64.zip && \
    unzip terraform_0.9.5_linux_amd64.zip -d /opt/terraform && \
    ln -s /opt/terraform/terraform /usr/bin/terraform && \
    # clean up
    apt-get --purge remove -y -q --auto-remove build-essential python-pip curl \
    unzip libssl-dev libffi-dev python-dev && \
    apt-get clean && rm -rf /var/lib/apt/lists/* && rm -rf /tmp/* && \
    # create project home directory and remove CloudCore configuration
    mkdir -p "/data" && \
    echo 'localhost' > /etc/ansible/hosts && \
    rm /etc/cloud-core.conf && rm /etc/krb5.conf


ENV CLOUD_CORE_HOST="0.0.0.0" CLOUD_CORE_PORT=80 PROJECTS_HOME="/data" \
    CLOUD_CORE_DB_USER=cloudcore CLOUD_CORE_DB_PASSWORD=cloudcore CLOUD_CORE_DB_NAME=cloudcoredb \
    CLOUD_CORE_DB_REPLICA="" CLOUD_CORE_DB_HOSTS="mongo:27017" KRB5_CONFIG="/data/krb5.conf" \
    ANSIBLE_CONFIG="/etc/ansible/ansible.cfg" CLOUD_CORE_RABBITMQ="amqp://cloudcore:cloudcore@localhost:5672/cloudcore"

USER cloudcore

VOLUME /data

WORKDIR /data

CMD ["cloudcored"]
