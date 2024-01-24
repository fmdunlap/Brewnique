import { env } from '$env/dynamic/private';
import postmark from 'postmark';

const client = new postmark.ServerClient(env.POSTMARK_API_KEY);

export async function sendEmail(to: string, subject: string, html: string, text?: string) {
	return client.sendEmail({
		From: 'no-reply@brewnique.io',
		To: to,
		Subject: subject,
		HtmlBody: html,
		TextBody: text
	});
}
