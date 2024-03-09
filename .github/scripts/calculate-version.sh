#!/bin/bash

calculate_version() {
    current_tag=$1
    commit_messages=$2

    current_version=$(echo "${current_tag#v}")

    major=$(echo "$current_version" | cut -d. -f1)
    minor=$(echo "$current_version" | cut -d. -f2)
    patch=$(echo "$current_version" | cut -d. -f3)

    for message in $commit_messages; do
        if [[ $message == *"!:"* ]]; then
            ((major++))
        elif [[ $message == *"feat:"* ]]; then
            ((minor++))
        elif [[ $message == *"fix:"* ]]; then
            ((patch++))
        fi
    done

    echo "v$major.$minor.$patch"
}

current_tag=$(git describe --tags --abbrev=0 2>/dev/null || echo "v-1.0.0")

case $current_tag in
    "v-1.0.0") commit_messages=$(git log --pretty=format:"%s") ;;
    *) commit_messages=$(git log --pretty=format:"%s" $current_tag...HEAD)
esac

version=$(calculate_version "$current_tag" "$commit_messages")

echo $version