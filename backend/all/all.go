package all

import (
	// Active file systems
	_ "github.com/rclone/rclone/backend/drive"
	_ "github.com/rclone/rclone/backend/googlecloudstorage"
	_ "github.com/rclone/rclone/backend/googlephotos"
	_ "github.com/rclone/rclone/backend/onedrive"
)
