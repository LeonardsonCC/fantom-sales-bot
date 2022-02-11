import { CacheType, CommandInteraction } from "discord.js";
import { CommandFunction } from ".";

const command: CommandFunction = async (
  interaction: CommandInteraction<CacheType>
) => {
  return await interaction.reply("pong");
};

export default command;
