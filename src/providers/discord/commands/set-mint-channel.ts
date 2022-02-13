import { CacheType, Client, CommandInteraction } from "discord.js";
import { CommandFunction } from ".";
import { isValidAddress } from "../../blockchain";
import * as Database from "../../database";

const command: CommandFunction = async (
  interaction: CommandInteraction<CacheType>,
  client: Client
) => {
  const channelId = interaction.options.getChannel("channel")?.id;
  const collectionAddr = interaction.options.getString("collection");
  if (channelId && collectionAddr !== null) {
    if (!isValidAddress(collectionAddr))
      return interaction.reply("Sorry, the collection address is invalid");

    console.log(
      "SETTING MINT CHANNEL TO SERVER: ",
      interaction.guildId,
      channelId,
      collectionAddr
    );
    const channel = client.channels.cache.get(channelId);
    if (channel && channel.isText()) {
      try {
        await channel.send("The mints will appear here!");
      } catch (err) {
        await interaction.reply("I can't post on that channel...");
        return;
      }
      const db = Database.connect();

      try {
        db.collection("mint").updateOne(
          {
            collection: collectionAddr,
          },
          {
            $set: {
              [`guilds.${interaction.guildId!.toString()}`]: {
                mintChannelId: channelId,
              },
            },
          },
          {
            upsert: true,
          }
        );
      } catch (err) {
        console.log("ERROR UPDATING...", err);
        db.collection("mint").insertOne({
          collection: collectionAddr,
          [`guilds.${interaction.guildId!.toString()}`]: {
            mintChannelId: channelId,
          },
        });
      } finally {
        interaction.reply("Configuration done!");
      }
    }
  }
};

export default command;
