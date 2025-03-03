import type { CustomThemeConfig } from '@skeletonlabs/tw-plugin';

export const aseTheme: CustomThemeConfig = {
	name: 'ase-custom-theme',
	properties: {
		// =~= Theme Properties =~=
		'--theme-font-family-base': `Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, 'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji'`,
		'--theme-font-family-heading': `system-ui`,
		'--theme-font-color-base': '0 0 0',
		'--theme-font-color-dark': '255 255 255',
		'--theme-rounded-base': '0px',
		'--theme-rounded-container': '0px',
		'--theme-border-base': '1px',
		// =~= Theme On-X Colors =~=
		'--on-primary': '255 255 255',
		'--on-secondary': '0 0 0',
		'--on-tertiary': '255 255 255',
		'--on-success': '255 255 255',
		'--on-warning': '0 0 0',
		'--on-error': '255 255 255',
		'--on-surface': '255 255 255',
		// =~= Theme Colors  =~=
		// primary | #0089d9
		'--color-primary-50': '217 237 249', // #d9edf9
		'--color-primary-100': '204 231 247', // #cce7f7
		'--color-primary-200': '191 226 246', // #bfe2f6
		'--color-primary-300': '153 208 240', // #99d0f0
		'--color-primary-400': '77 172 228', // #4dace4
		'--color-primary-500': '0 137 217', // #0089d9
		'--color-primary-600': '0 123 195', // #007bc3
		'--color-primary-700': '0 103 163', // #0067a3
		'--color-primary-800': '0 82 130', // #005282
		'--color-primary-900': '0 67 106', // #00436a
		// secondary | #E4DFDA
		'--color-secondary-50': '251 250 249', // #fbfaf9
		'--color-secondary-100': '250 249 248', // #faf9f8
		'--color-secondary-200': '248 247 246', // #f8f7f6
		'--color-secondary-300': '244 242 240', // #f4f2f0
		'--color-secondary-400': '236 233 229', // #ece9e5
		'--color-secondary-500': '228 223 218', // #E4DFDA
		'--color-secondary-600': '205 201 196', // #cdc9c4
		'--color-secondary-700': '171 167 164', // #aba7a4
		'--color-secondary-800': '137 134 131', // #898683
		'--color-secondary-900': '112 109 107', // #706d6b
		// tertiary | #1E3231
		'--color-tertiary-50': '221 224 224', // #dde0e0
		'--color-tertiary-100': '210 214 214', // #d2d6d6
		'--color-tertiary-200': '199 204 204', // #c7cccc
		'--color-tertiary-300': '165 173 173', // #a5adad
		'--color-tertiary-400': '98 112 111', // #62706f
		'--color-tertiary-500': '30 50 49', // #1E3231
		'--color-tertiary-600': '27 45 44', // #1b2d2c
		'--color-tertiary-700': '23 38 37', // #172625
		'--color-tertiary-800': '18 30 29', // #121e1d
		'--color-tertiary-900': '15 25 24', // #0f1918
		// success | #459615
		'--color-success-50': '227 239 220', // #e3efdc
		'--color-success-100': '218 234 208', // #daead0
		'--color-success-200': '209 229 197', // #d1e5c5
		'--color-success-300': '181 213 161', // #b5d5a1
		'--color-success-400': '125 182 91', // #7db65b
		'--color-success-500': '69 150 21', // #459615
		'--color-success-600': '62 135 19', // #3e8713
		'--color-success-700': '52 113 16', // #347110
		'--color-success-800': '41 90 13', // #295a0d
		'--color-success-900': '34 74 10', // #224a0a
		// warning | #ee942e
		'--color-warning-50': '252 239 224', // #fcefe0
		'--color-warning-100': '252 234 213', // #fcead5
		'--color-warning-200': '251 228 203', // #fbe4cb
		'--color-warning-300': '248 212 171', // #f8d4ab
		'--color-warning-400': '243 180 109', // #f3b46d
		'--color-warning-500': '238 148 46', // #ee942e
		'--color-warning-600': '214 133 41', // #d68529
		'--color-warning-700': '179 111 35', // #b36f23
		'--color-warning-800': '143 89 28', // #8f591c
		'--color-warning-900': '117 73 23', // #754917
		// error | #EE2E31
		'--color-error-50': '252 224 224', // #fce0e0
		'--color-error-100': '252 213 214', // #fcd5d6
		'--color-error-200': '251 203 204', // #fbcbcc
		'--color-error-300': '248 171 173', // #f8abad
		'--color-error-400': '243 109 111', // #f36d6f
		'--color-error-500': '238 46 49', // #EE2E31
		'--color-error-600': '214 41 44', // #d6292c
		'--color-error-700': '179 35 37', // #b32325
		'--color-error-800': '143 28 29', // #8f1c1d
		'--color-error-900': '117 23 24', // #751718
		// surface | #1B1B1B
		'--color-surface-50': '221 221 221', // #dddddd
		'--color-surface-100': '209 209 209', // #d1d1d1
		'--color-surface-200': '198 198 198', // #c6c6c6
		'--color-surface-300': '164 164 164', // #a4a4a4
		'--color-surface-400': '95 95 95', // #5f5f5f
		'--color-surface-500': '27 27 27', // #1B1B1B
		'--color-surface-600': '24 24 24', // #181818
		'--color-surface-700': '20 20 20', // #141414
		'--color-surface-800': '16 16 16', // #101010
		'--color-surface-900': '13 13 13' // #0d0d0d
	}
};
