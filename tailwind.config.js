/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/templates/*.gohtml",
    "./src/templates/**/*.gohtml",
    "./src/*css"
  ],
  theme: {
    extend: {},
  }
}

