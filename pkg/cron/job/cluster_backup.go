package job

import (
	"math"
	"sync"
	"time"

	"github.com/KubeOperator/KubeOperator/pkg/db"
	"github.com/KubeOperator/KubeOperator/pkg/dto"
	"github.com/KubeOperator/KubeOperator/pkg/logger"
	"github.com/KubeOperator/KubeOperator/pkg/model"
	"github.com/KubeOperator/KubeOperator/pkg/repository"
	"github.com/KubeOperator/KubeOperator/pkg/service"
)

type ClusterBackup struct {
	cLusterBackupFileService        service.CLusterBackupFileService
	clusterBackupStrategyRepository repository.ClusterBackupStrategyRepository
}

func NewClusterBackup() *ClusterBackup {
	return &ClusterBackup{
		cLusterBackupFileService:        service.NewClusterBackupFileService(),
		clusterBackupStrategyRepository: repository.NewClusterBackupStrategyRepository(),
	}
}

func (c *ClusterBackup) Run() {
	var wg sync.WaitGroup
	sem := make(chan struct{}, 2) // 信号量
	clusterBackupStrategies, _ := c.clusterBackupStrategyRepository.List()
	for _, clusterBackupStrategy := range clusterBackupStrategies {
		if clusterBackupStrategy.Status == "ENABLE" {
			var backupFiles []model.ClusterBackupFile
			db.DB.Find(&backupFiles)
			if len(backupFiles) < clusterBackupStrategy.SaveNum {
				if len(backupFiles) > 0 {
					lastBackupFile := backupFiles[0]
					backupDate := lastBackupFile.CreatedAt
					now := time.Now()
					sumD := now.Sub(backupDate)
					day := int(math.Floor(sumD.Hours() / 24))
					if day < clusterBackupStrategy.Cron {
						continue
					}
				}
				wg.Add(1)
				go func() {
					defer wg.Done()
					sem <- struct{}{}
					defer func() { <-sem }()
					var cluster model.Cluster
					db.DB.Where("id = ?", clusterBackupStrategy.ClusterID).Find(&cluster)
					logger.Log.Infof("start backup cluster [%s]", cluster.Name)
					if cluster.ID != "" {
						err := c.cLusterBackupFileService.Backup(dto.ClusterBackupFileCreate{ClusterName: cluster.Name})
						if err != nil {
							logger.Log.Errorf("backup cluster error: %s", err.Error())
						}
					}
				}()
			}
		}
	}
	wg.Wait()
}
