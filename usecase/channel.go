package usecase

type ChannelUseCase interface {
	func GetRecommendChannels()
	func GetChannelByUserID(id: string)
	func CreateChannel()
}

type ChannelBiz struct {
	repository: 
}
func NewChannelUseCase() -> ChannelUseCase {
	
}
