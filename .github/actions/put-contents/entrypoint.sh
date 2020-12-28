#!/bin/sh
#commit file to the destination branch

#test
mkdir -p delete/me
FILE_TO_COMMIT="delete/me/delete_me.txt"
date >> $FILE_TO_COMMIT
cat $FILE_TO_COMMIT

export MESSAGE="generated $FILE_TO_COMMIT"
export SHA=$(git rev-parse $DESTINATION_BRANCH:$FILE_TO_COMMIT)
export CONTENT=$(base64 $FILE_TO_COMMIT)
echo "$DESTINATION_BRANCH:$FILE_TO_COMMIT:$SHA"

if [ "$SHA" = "$DESTINATION_BRANCH:$FILE_TO_COMMIT" ]; then
    echo "not exist"
    gh api --method PUT /repos/:owner/:repo/contents/$FILE_TO_COMMIT \
        --field message="$MESSAGE" \
        --field content="$CONTENT" \
        --field encoding="base64" \
        --field branch="$DESTINATION_BRANCH"
else
    echo "exist"
    gh api --method PUT /repos/:owner/:repo/contents/$FILE_TO_COMMIT \
        --field message="$MESSAGE" \
        --field content="$CONTENT" \
        --field encoding="base64" \
        --field branch="$DESTINATION_BRANCH" \
        --field sha="$SHA"
fi
