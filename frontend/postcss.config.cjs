/** @type {import('postcss').Config} */
module.exports = {
  plugins: {
    "@tailwindcss/postcss": {},   // ← これだけ残せば OK
  },
};
