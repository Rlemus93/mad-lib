/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.html", "./public/**/*.js"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["light", "dark"], // Add themes you want to use
  },
}
