#!/bin/sh

STAGE=$1

if [ -z "$STAGE" ]; then
    echo "❌ Error: No stage parameter provided! Usage: deploy.sh [dev|prod]"
    exit 1
fi

case "$STAGE" in
    dev)
        # Create S3
        # Create Cloudfront
        # Create EC2 
        # Compile go executable
        # Copy go executable to EC2
        # Build npm
        # Copy npm build to S3  
        ;;
    prod)
        echo "🚀 Starting PROD deployment pipeline..."
        ;;
    *)
        echo "❓ Unknown stage: '$STAGE'. Please use 'dev' or 'prod'."
        exit 1
        ;;
esac