export function convertBase64ToFile(base64: string) {
	const byteString = atob(base64.split(',')[1]);
	const mimeString = base64.split(',')[0].split(':')[1].split(';')[0];
	const ab = new ArrayBuffer(byteString.length);
	const ia = new Uint8Array(ab);
	for (let i = 0; i < byteString.length; i++) ia[i] = byteString.charCodeAt(i);
	const blob = new Blob([ab], { type: mimeString });
	return new File([blob], 'image', { type: mimeString });
}
