package marketplaces

type IMarketplace interface {
	Subscribe(chan *Sale) error
}
