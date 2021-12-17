import { TwitterApi } from "twitter-api-v2";
import dotenv from "dotenv";

dotenv.config();

const client = new TwitterApi({
  appKey: process.env.CONSUMER_KEY!,
  appSecret: process.env.CONSUMER_SECRET!,
  accessToken: process.env.ACCESS_TOKEN!,
  accessSecret: process.env.ACCESS_SECRET!,
});

const tweet = async (message: string, image?: string) => {
  // client.v2.tweet(message).then(console.log).catch(console.log);
  if (image) {
    try {
      const mediaId = await client.v1.uploadMedia(image, { type: "PNG" });
      await client.v2.tweet(message, {
        media: {
          media_ids: [mediaId],
        },
      });
    } catch (err) {
      await client.v2.tweet(message);
    }
  } else {
    console.log("Image doesn't work");
    await client.v2.tweet(message);
  }
};

export { tweet };
