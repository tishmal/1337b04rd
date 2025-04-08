package port

type AvatarService interface {
	GetRandomAvatar() (name string, avatarURL string, err error)
	GetTotalAvatarCount() (int, error)
}
