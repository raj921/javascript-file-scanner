// tailwind.config.js
module.exports = {
  content: [
    './web/**/*.html',
    './web/**/*.js',
    './web/**/*.jsx',
    './web/**/*.ts',
    './web/**/*.tsx',
  ],
  theme: {
    extend: {
      colors: {
        background: 'var(--background)',
        foreground: 'var(--foreground)',
        card: 'var(--card)',
        // Add other custom colors here
      },
      // Define other customizations here
    },
  },
  plugins: [],
}


