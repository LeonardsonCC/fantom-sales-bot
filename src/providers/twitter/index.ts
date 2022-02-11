import { TwitterApi } from "twitter-api-v2";
import dotenv from "dotenv";
import { Sale } from "../../types/sale";
import { Action, makeMessage } from "./message";

dotenv.config();

const client = new TwitterApi({
  appKey: process.env.CONSUMER_KEY!,
  appSecret: process.env.CONSUMER_SECRET!,
  accessToken: process.env.ACCESS_TOKEN!,
  accessSecret: process.env.ACCESS_SECRET!,
});

const post = async (
  sale: Sale,
  actionBefore: Sale,
  action: Action,
  imageUrl?: string,
  image?: string,
  imageType: "PNG" | "JPG" | "GIF" = "PNG"
) => {
  console.log("IMAGE TYPE", imageType);

  const message = await makeMessage(sale, actionBefore, action);

  if (image) {
    try {
      const mediaId = await client.v1.uploadMedia(image, { type: imageType });
      await client.v2.tweet(message, {
        media: {
          media_ids: [mediaId],
        },
      });
    } catch (err) {
      console.log("error uploading image... ", err);
      await client.v2.tweet(message);
    }
  } else {
    console.log("image not sent... ");
    await client.v2.tweet(message);
  }
};

export { post };
