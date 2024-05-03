
/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
			animation: {
				fadeIn: 'fadeIn .3s ease-in-out',
				fadeOut: 'fadeOut ease-in-out',
				contract: 'contract linear',
				flash: 'flash 1s linear',
			},

			keyframes: {
				fadeIn: {
					from: { opacity: 0 },
					to: { opacity: 1 },
				},
				fadeOut: {
					"0%": { opacity: 0, transform: "translateX(0)"},
					"20%": { opacity: 1, transform: "translateX(0)" },
					"90%": { opacity: 1, transform: "translateX(0)" },
					"100%": { opacity: 0 , transform: "translateX(-100%)"},
				},
				contract: {
					"0%": { width: "100%" },
					"100%": { width: "0%" },
				},
				flash: {
					"0%, 100%": { opacity: 1 },
					"50%": { opacity: 0 },
				},
			},
		},
  },
  plugins: [],
}

