#!/usr/bin/env bash

set -ueo pipefail
set +x

cleanup(){
    for version in ${unique_versions}; do
        git worktree remove --force $(version_dir $version)
    done
    rm -rf "$TMP"
    echo "cleaned up test successfully"
}
#trap cleanup EXIT

# TODO make sure the number of storagenode versions matches the number of sns from setup

stage1_sat_version="$1"
stage1_uplink_versions="$2"
stage1_storagenode_versions="$3"
stage2_sat_version="$4"
stage2_uplink_versions="$5"
stage2_storagenode_versions="$6"

echo "stage1_sat_version" $stage1_sat_version
echo "stage1_uplink_versions" $stage1_uplink_versions
echo "stage1_storagenode_versions" $stage1_storagenode_versions
echo "stage2_sat_version" $stage2_sat_version
echo "stage2_uplink_versions" $stage2_uplink_versions
echo "stage2_storagenode_versions" $stage2_storagenode_versions

TMP=$(mktemp -d -t tmp.XXXXXXXXXX)

find_unique_versions(){
    echo "$*" | tr " " "\n" | sort | uniq
}

version_dir(){
    echo "${TMP}/${1}"
}

replace_in_file(){
    local src="$1"
    local dest="$2"
    local path=$3
    sed -i '' "s#${src}#${dest}#g" "${path}"
}

setup_stage(){
    local test_dir=$1
    local sat_version=$2
    local stage_sn_versions=$3

    echo "Satellite version: ${sat_version}"
    echo "Storagenode versions: ${stage_sn_versions}"

    local src_sat_version_dir=$(version_dir ${sat_version})

    PATH=$src_sat_version_dir/bin:$PATH src_sat_cfg_dir=$(storj-sim network env --config-dir=${src_sat_version_dir}/local-network/ SATELLITE_0_DIR)
    PATH=$test_dir/bin:$PATH dest_sat_cfg_dir=$(storj-sim network env --config-dir=${test_dir}/local-network/ SATELLITE_0_DIR)

    # ln binary and copy config.yaml for desired version
    ln -f $src_sat_version_dir/bin/satellite $dest_sat_cfg_dir/satellite
    cp $src_sat_cfg_dir/config.yaml $dest_sat_cfg_dir
    replace_in_file "${src_sat_version_dir}" "${test_dir}" "${dest_sat_cfg_dir}/config.yaml"

    counter=0
    for sn_version in ${stage_sn_versions}; do
        local src_sn_version_dir=$(version_dir ${sn_version})

        PATH=$src_sn_version_dir/bin:$PATH src_sn_cfg_dir=$(storj-sim network env --config-dir=${src_sn_version_dir}/local-network/ STORAGENODE_${counter}_DIR)
        PATH=$test_dir/bin:$PATH dest_sn_cfg_dir=$(storj-sim network env --config-dir=${test_dir}/local-network/ STORAGENODE_${counter}_DIR)

        dest_sat_nodeid=$(grep "storage.whitelisted-satellites" ${dest_sn_cfg_dir}/config.yaml)
        src_sat_nodeid=$(grep "storage.whitelisted-satellites" "${src_sn_cfg_dir}/config.yaml")

        # ln binary and copy config.yaml for desired version
        ln -f $src_sn_version_dir/bin/storagenode $dest_sn_cfg_dir/storagenode
        cp $src_sn_cfg_dir/config.yaml $dest_sn_cfg_dir

        # update config dir in config.yaml as well as whitelisted satellites in config.yaml
        replace_in_file "${src_sn_version_dir}" "${test_dir}" "${dest_sn_cfg_dir}/config.yaml"
        replace_in_file  "${src_sat_nodeid}" "${dest_sat_nodeid}" "${dest_sn_cfg_dir}/config.yaml"

        let counter+=1
    done

#    for ul_version in ${stage_ul_versions}; do
        # use desired uplink binary and config
 #   done
    # PATH=$src_ul_version_dir/bin:$PATH src_ul_cfg_dir=$(storj-sim network env --config-dir=${src_ul_version_dir}/local-network/ GATEWAY_0_DIR)
    # PATH=$test_dir/bin:$PATH dest_ul_cfg_dir=$(storj-sim network env --config-dir=${test_dir}/local-network/ GATEWAY_0_DIR)

    # src_ul_scope=$(grep "scope" "${src_ul_cfg_dir}/config.yaml")
    # dest_ul_scope=$(grep "scope" "${dest_ul_cfg_dir}/config.yaml")

    # cp $src_ul_cfg_dir/config.yaml $dest_ul_cfg_dir
    # replace_in_file "${src_ul_version_dir}" "${test_dir}" "${dest_ul_cfg_dir}/config.yaml"
    # replace_in_file "${src_ul_scope}" "${dest_ul_scope}" "${dest_ul_cfg_dir}/config.yaml"
}

# Set up each environment
unique_versions=$(find_unique_versions "$stage1_sat_version" "$stage1_uplink_versions" "$stage1_storagenode_versions" "$stage2_sat_version" "$stage2_uplink_versions" "$stage2_storagenode_versions")

STORJ_NETWORK_HOST4=${STORJ_NETWORK_HOST4:-127.0.0.1}
STORJ_SIM_POSTGRES=${STORJ_SIM_POSTGRES:-""}

if [ -z ${STORJ_SIM_POSTGRES} ]; then
    echo "STORJ_SIM_POSTGRES is required for the satellite DB. Example: STORJ_SIM_POSTGRES=postgres://[user]:[pass]@[host]/[db]?sslmode=disable"
    exit 1
fi

echo "Setting up environments for versions" ${unique_versions}

# Get latest release tags and clean up git worktree
git fetch --tags
git worktree prune
for version in ${unique_versions}; do
    dir=$(version_dir ${version})
    bin_dir=${dir}/bin

    echo -e "\nAdding worktree for ${version} in ${dir}."
    git worktree add -f ${dir} ${version}
    rm ${dir}/internal/version/release.go
    echo "Installing storj-sim for ${version} in ${dir}."
    GOBIN=${bin_dir} make -C "${dir}" install-sim > /dev/null 2>&1
    echo "Setting up storj-sim for ${version}. Bin: ${bin_dir}, Config: ${dir}/local-network"
    PATH=${bin_dir}:$PATH storj-sim -x --host="${STORJ_NETWORK_HOST4}" --postgres="${STORJ_SIM_POSTGRES}" --config-dir "${dir}/local-network" network setup > /dev/null 2>&1

    echo "Finished setting up. ${dir}/local-network:" $(ls ${dir}/local-network)
    echo "Binary shasums:"
    shasum ${bin_dir}/satellite
    shasum ${bin_dir}/storagenode
    shasum ${bin_dir}/uplink
done

# Use stage 1 satellite version as the starting state. Create a cp of that
# version folder so we don't worry about dirty states. Then copy/link/mv
# appropriate resources into that folder to ensure we have correct versions.
test_dir=$(version_dir "test_dir")
cp -r $(version_dir ${stage1_sat_version}) ${test_dir}
echo -e "\nSetting up stage 1 in ${test_dir}"
scriptdir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
setup_stage "${test_dir}" "${stage1_sat_version}" "${stage1_storagenode_versions}"
for ul_version in ${stage1_uplink_versions}; do
    echo "Uplink version: ${ul_version}"
    src_ul_version_dir=$(version_dir ${ul_version})
    ln -f ${src_ul_version_dir}/bin/uplink $test_dir/bin/uplink
    PATH=$test_dir/bin:$PATH storj-sim -x --host "${STORJ_NETWORK_HOST4}" --config-dir "${test_dir}/local-network" network test bash "${scriptdir}/test-versions.sh" "${test_dir}/local-network" "upload" "${ul_version}"
done


echo -e "\nSetting up stage 2 in ${test_dir}"
setup_stage "${test_dir}" "${stage2_sat_version}" "${stage2_storagenode_versions}"
echo -e "\nRunning stage 2."
for ul_version in ${stage2_uplink_versions}; do
    echo "Uplink version: ${ul_version}"
    src_ul_version_dir=$(version_dir ${ul_version})
    ln -f ${src_ul_version_dir}/bin/uplink $test_dir/bin/uplink
    PATH=$test_dir/bin:$PATH storj-sim -x --host "${STORJ_NETWORK_HOST4}" --config-dir "${test_dir}/local-network" network test bash "${scriptdir}/test-versions.sh" "${test_dir}/local-network" "download" "${stage1_uplink_versions}"
done


echo -e "\nCleaning up."
PATH=$test_dir/bin:$PATH storj-sim -x --host "${STORJ_NETWORK_HOST4}" --config-dir "${test_dir}/local-network" network test bash "${scriptdir}/test-versions.sh" "${test_dir}/local-network" "cleanup" "${stage1_uplink_versions}"
