import { CacheType, Client, CommandInteraction } from "discord.js";
import { CommandFunction } from ".";
import * as Database from "../../database";

const command: CommandFunction = async (
  interaction: CommandInteraction<CacheType>,
  client: Client
) => {
  const channelId = interaction.options.getChannel("channel")?.id;
  if (channelId) {
    console.log("SETTING CHANNEL TO SERVER: ", interaction.guildId, channelId);

    const channel = client.channels.cache.get(channelId);
    if (channel && channel.isText()) {
      try {
        await channel.send("The sales will appear here!");
      } catch (err) {
        await interaction.user.send("I can't post on that channel...");
        return;
      }
      const db = Database.connect();

      try {
        db.collection("bot").updateOne(
          {
            guildId: interaction.guildId,
          },
          {
            $set: { salesChannelId: channel.id },
          },
          {
            upsert: true,
          }
        );
      } catch (err) {
        console.log("ERROR UPDATING...", err);
        db.collection("bot").insertOne({
          guildId: interaction.guildId,
          salesChannelId: channel.id,
        });
      }
    }
  }
};

export default command;
