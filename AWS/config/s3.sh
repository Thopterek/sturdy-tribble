#!/bin/sh

filename="$1"
source "./Logs/logger.sh"

echo "=> [.][🪣] Creating S3 Bucket in AWS...."

if [ -z "$BUCKET_NAME" ] || [ -z "$AWS_DEFAULT_REGION" ]; then
	echo "=> [❌][🪣] Failed to import one or more env parameters:"
	logger "ERROR" "Failed to import one or more env parameters:" "$filename" "-1"
	logger "ERROR" "Bucket name: '$BUCKET_NAME'." "$filename" "-1"
	logger "ERROR" "Default region: '$AWS_DEFAULT_REGION'." "$filename" "-1"
	exit 1
else
	echo "=> [✅][🪣] Env parameters imported successfully!"
	logger "INFO" "Env parameters imported successfully!" "$filename" "0"
fi

logger "INFO" "Deploying to region: '$AWS_DEFAULT_REGION'..." "$filename" "0"


aws s3api create-bucket \
    --bucket "$BUCKET_NAME" \
    --region "$AWS_DEFAULT_REGION" \
    --create-bucket-configuration LocationConstraint="$AWS_DEFAULT_REGION" \
    >/dev/null 2>&1

STATUS=$?

if [ "$STATUS" -eq 0 ]; then
	aws s3api put-public-access-block \
		--bucket "$BUCKET_NAME" \
		--public-access-block-configuration "BlockPublicAcls=true,IgnorePublicAcls=true,BlockPublicPolicy=true,RestrictPublicBuckets=true" \
		>/dev/null 2>&1

	aws s3api delete-bucket-encryption --bucket "$BUCKET_NAME" >/dev/null 2>&1

	aws s3api put-object --bucket "$BUCKET_NAME" --key "logs/"
	aws s3api put-object --bucket "$BUCKET_NAME" --key "data/"
	aws s3api put-object --bucket "$BUCKET_NAME" --key "uploads/"

    echo "=> [✅][🪣] Created new S3 Bucket: '$BUCKET_NAME'."
	logger "INFO" "Success! New S3 Bucket '$BUCKET_NAME' created from scratch." "$filename" "0"
else
    case "$STATUS" in
		255)
            echo "=> [ℹ️][🪣] '$BUCKET_NAME' already exists and you own it."
			logger "WARNING" "'$BUCKET_NAME' already exists and you own it." "$filename" "0"
            ;;
        *)
            echo "=> [❌][🪣] Error: Failed to create bucket: '$BUCKET_NAME'."
			logger "ERROR" "Failed to create bucket '$BUCKET_NAME':." "$filename" "-1"
			logger "ERROR" "Details: '$STATUS'" "$filename" "-1"
            exit 1
            ;;
    esac
fi
