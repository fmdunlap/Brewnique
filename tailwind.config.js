/** @type {import('tailwindcss').Config} */
const config = {
	darkMode: ['class'],
	content: [
		'./src/**/*.{html,js,svelte,ts}',
		'./node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}'
	],
	safelist: ['dark'],
	theme: {
		container: {
			center: true,
			padding: '2rem',
			screens: {
				'2xl': '1400px'
			}
		},
		extend: {
			colors: {
				primary: {
					50: '#fef7ee',
					100: '#fdedd7',
					200: '#fad8ae',
					300: '#f6bb7b',
					400: '#f19446',
					500: '#ed7621',
					600: '#de5d18',
					700: '#de5d18',
					800: '#933919',
					900: '#763118',
					950: '#40160a'
				},
				secondary: {
					50: '#f3f4f1',
					100: '#e5e7e0',
					200: '#ccd2c4',
					300: '#b9c1ae',
					400: '#909b80',
					500: '#727f63',
					600: '#59644c',
					700: '#464e3d',
					800: '#3a4034',
					900: '#34382f',
					950: '#1a1d16'
				},
				accent: {
					50: '#f4f6f3',
					100: '#e5e9e2',
					200: '#cad4c6',
					300: '#9bad94',
					400: '#7d9275',
					500: '#5b7455',
					600: '#455b40',
					700: '#364933',
					800: '#2d3b2a',
					900: '#253123',
					950: '#141b13'
				},
				background: {
					light: {
						primary: '#fff',
						secondary: '#f8fafc'
					},
					dark: {
						primary: '#020617',
						secondary: '#0f172a'
					}
				}
			}
		}
	},
	plugins: [require('flowbite/plugin')]
};

export default config;
