package payloads

import "github.com/denkylabs/discord-api-types-go/globals"

// https://discord.com/developers/docs/resources/sticker#sticker-object
type APISticker struct {
	// ID of the sticker
	Id globals.Snowflake `json:"id"`
	// For standard stickers, ID of the pack the sticker is from
	PackId globals.Snowflake `json:"pack_id"`
	// Name of the sticker
	Name string `json:"name"`
	// Description of the sticker
	Description string `json:"description"`
	// For guild stickers, the Discord name of a unicode emoji representing the sticker's expression. for standard stickers, a comma-separated list of related expressions.
	Tags string `json:"tags"`
	// Previously the sticker asset hash, now an empty string
	Asset string `json:"asset"`
	// Type of sticker
	// See https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-types
	Type StickerType `json:"type"`
	// Type of sticker format
	// See https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-format-types
	FormatType StickerFormatType `json:"format"`
	// Whether this guild sticker can be used, may be false due to loss of Server Boosts
	Available bool `json:"available"`
	// ID of the guild that owns this sticker
	GuildId globals.Snowflake `json:"guild_id"`
	// The user that uploaded the guild sticker
	User APIUser `json:"user"`
	// The standard sticker's sort order within its pack
	SortValue uint16 `json:"sort_value"`
}

// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-types
type StickerType uint8

// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-types
const (
	// An official sticker in a pack, part of Nitro or in a removed purchasable pack
	StickerTypeStandard StickerType = 1
	// A sticker uploaded to a Boosted guild for the guild's members
	StickerTypeGuild StickerType = 2
)

// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-format-types
type StickerFormatType uint8

// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-format-types
const (
	StickerFormatTypePNG    StickerFormatType = 1
	StickerFormatTypeAPNG   StickerFormatType = 2
	StickerFormatTypeLottie StickerFormatType = 3
)

// https://discord.com/developers/docs/resources/sticker#sticker-item-object
type APIStickerItem struct {
	// ID of the sticker
	Id globals.Snowflake `json:"id"`
	// Name of the sticker
	Name string `json:"name"`
	// Type of sticker format
	// See https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-format-types
	FormatType StickerFormatType `json:"format"`
}

// https://discord.com/developers/docs/resources/sticker#sticker-pack-object
type APIStickerPack struct {
	// ID of the sticker pack
	Id globals.Snowflake `json:"id"`
	// The stickers in the pack
	Stickers []APISticker `json:"stickers"`
	// Name of the sticker pack
	Name string `json:"name"`
	// ID of the pack's SKU
	SkuId globals.Snowflake `json:"sku_id"`
	// ID of a sticker in the pack which is shown as the pack's icon
	CoverStickerId globals.Snowflake `json:"cover_sticker_id"`
	// Description of the sticker pack
	Description string `json:"description"`
	// ID of the sticker pack's banner image
	BannerAssetId globals.Snowflake `json:"banner_asset_id"`
}
