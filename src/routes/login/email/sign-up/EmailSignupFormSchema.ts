import { z } from 'zod';

export const emailSignupFormSchema = z
	.object({
		email: z.string().email(),
		password: z.string().min(8).max(100),
		confirmPassword: z.string().min(8).max(100)
	})
	.superRefine(({ password }, checkPassComplexity) => {
		const containsUppercase = (ch: string) => /[A-Z]/.test(ch);
		const containsLowercase = (ch: string) => /[a-z]/.test(ch);
		const containsSpecialChar = (ch: string) => /[`!@#$%^&*()_\-+=[\]{};':"\\|,.<>/?~ ]/.test(ch);
		let countOfUpperCase = 0,
			countOfLowerCase = 0,
			countOfNumbers = 0,
			countOfSpecialChar = 0;
		for (let i = 0; i < password.length; i++) {
			const ch = password.charAt(i);
			if (!isNaN(+ch)) countOfNumbers++;
			else if (containsUppercase(ch)) countOfUpperCase++;
			else if (containsLowercase(ch)) countOfLowerCase++;
			else if (containsSpecialChar(ch)) countOfSpecialChar++;
		}
		if (countOfLowerCase < 1) {
			checkPassComplexity.addIssue({
				code: 'custom',
				message: 'Password must have at least one lowercase letter.'
			});
		}
		if (countOfUpperCase < 1) {
			checkPassComplexity.addIssue({
				code: 'custom',
				message: 'Password must have at least one uppercase letter.'
			});
		}
		if (countOfNumbers < 1) {
			checkPassComplexity.addIssue({
				code: 'custom',
				message: 'Password must have at least one number.'
			});
		}
		if (countOfSpecialChar < 1) {
			checkPassComplexity.addIssue({
				code: 'custom',
				message: 'Password must have at least one special character.'
			});
		}
	});
export type EmailSignupFormSchema = typeof emailSignupFormSchema;
