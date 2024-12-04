#!/usr/bin/env bash
# SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
# SPDX-License-Identifier: Apache-2.0

set -euo pipefail

mkdir -p /var/spool/slurmctld/
chown slurm:slurm /var/spool/slurmctld/
chmod 700 /var/spool/slurmctld/

mkdir -p /run/slurmctld/
chown slurm:slurm /run/slurmctld/

mkdir -p /etc/slurm/
cp /opt/etc/slurm/* /etc/slurm/
chown slurm:slurm /etc/slurm/*
chmod 644 /etc/slurm/*.conf
chmod 600 /etc/slurm/*.key

slurmctld -D
