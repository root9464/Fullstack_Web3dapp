import { Bot, InlineKeyboard } from 'grammy';

// Create an instance of the `Bot` class and pass your bot token to it.
const bot = new Bot('7904526434:AAFhRzmcViUvFEH2zORDhWVTf8GMKZCRCpE'); // <-- put your bot token between the ""

// You can now register listeners on your bot object `bot`.
// grammY will call the listeners when users send messages to your bot.

// Handle the /start command.
const inlineKeyboard = new InlineKeyboard().webApp('Earn Ton', 'https://anytapton.ru').url('Channel', 'https://t.me/stake_t');
// Make it interactive
const msg = `
Welcome to Earn Ton!
We’re happy to see you join our community of investors. Here, every TON can generate profit, even if you’re just getting started.

📲 Don’t miss out on exclusive insights and updates!
Make sure to subscribe to our Telegram channel to stay informed and get the latest strategies to maximize your earnings.

We’re here to help you make your money work smarter.

`;
// Handle other messages.
bot.command('start', async (ctx) => {
  await ctx.reply(msg, { reply_markup: inlineKeyboard });
});

// Now that you specified how to handle messages, you can start your bot.
// This will connect to the Telegram servers and wait for messages.

// Start the bot.
bot.start();
