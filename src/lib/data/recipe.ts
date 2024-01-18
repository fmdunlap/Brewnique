import { v4 as uuidv4 } from 'uuid';
import { S3 } from './s3';
import { PutObjectCommand } from '@aws-sdk/client-s3';

export async function uploadRecipePhoto(file: File, recipeId: string) {
	const filetype = file.type.split('/')[1];
	const key = `recipes/${recipeId}/${uuidv4()}.${filetype}`;
	console.log('uploading to ' + key);
	S3.send(
		new PutObjectCommand({
			Bucket: 'brewnique',
			Key: key,
			Body: Buffer.from(await file.arrayBuffer()),
			ContentType: file.type,
			ContentLength: file.size
		})
	);
	return 'https://cdn.brewnique.io/' + key;
}
