/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.{html,js,amber}"],
  theme: {
    extend: {
      fontFamily: {
        sans: ['"Inter"'],
        display: ['"Manrope"'],
        body: ['"Noto Sans"'],
      },
      colors: {
        dark: "#1B1B1B",
        light: "#FFF",
        accent: "#7B00D3",
        accentDark: "#FFDB4D",
        gray: "#747474",
      },
    },
  },
  plugins: [],
};
