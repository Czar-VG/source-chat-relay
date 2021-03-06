package bot

import (
	"bytes"
	"regexp"

	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
)

var ChannelRegex = regexp.MustCompile("^(?:<#)?([0-9]+)>?$")

func GuildMemberPermissions(member *discordgo.Member, guild *discordgo.Guild) (apermissions int) {
	if member.User.ID == guild.OwnerID {
		apermissions = discordgo.PermissionAll
		return
	}

	for _, role := range guild.Roles {
		if role.ID == guild.ID {
			apermissions |= role.Permissions
			break
		}
	}

	for _, role := range guild.Roles {
		for _, roleID := range member.Roles {
			if role.ID == roleID {
				apermissions |= role.Permissions
				break
			}
		}
	}

	if apermissions&discordgo.PermissionAdministrator != 0 {
		apermissions |= discordgo.PermissionAllChannel
	}

	return
}

func GetMessageGuild(c *exrouter.Context, m *discordgo.Message) (*discordgo.Guild, error) {
	channel, err := c.Channel(m.ChannelID)

	if err != nil {
		return nil, err
	}

	guild, err := c.Guild(channel.GuildID)

	if err != nil {
		return nil, err
	}

	return guild, nil
}

func CapitalChannelName(c *discordgo.Channel) string {
	nameBytes := []byte(c.Name)

	return string(bytes.ToUpper(nameBytes[:1])) + string(nameBytes[1:])
}

func ParseChannel(arg string) (string, bool) {
	if ChannelRegex.Match([]byte(arg)) {
		return ChannelRegex.FindAllStringSubmatch(arg, -1)[0][1], true
	}

	return "", false
}
