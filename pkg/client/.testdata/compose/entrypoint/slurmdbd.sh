#!/usr/bin/env bash
# SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
# SPDX-License-Identifier: Apache-2.0

set -euo pipefail

mkdir -p /run/slurmdbd/
chown slurm:slurm /run/slurmdbd/

mkdir -p /etc/slurm/
cp /opt/etc/slurm/* /etc/slurm/
chown slurm:slurm /etc/slurm/*
chmod 644 /etc/slurm/*.conf
chmod 600 /etc/slurm/slurmdbd.conf
chmod 600 /etc/slurm/*.key

slurmdbd -D
