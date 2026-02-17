/**
 * Mini Style Guide – colour & typography tokens
 * Import these wherever you need raw values outside of Tailwind.
 */

export const colors = {
  primary: '#F23568',
  secondary: '#056CF2',
  tertiary: '#48D9CA',
  accent: '#88BF11',
  error: '#F27B13',
  grey: {
    bg: '#f5f5f5',
    border: '#ddd',
  },
  alert: {
    success: { bg: '#e6f4ea', text: '#288c2a' },
    error: { bg: '#fff4e6', text: '#c23e00' },
  },
} as const

export const typography = {
  heading: {
    family: "'Montserrat', sans-serif",
    weight: 700,
    h1: '36px',
    h2: '30px',
    h3: '24px',
  },
  body: {
    family: "'Open Sans', sans-serif",
    size: '16px',
  },
} as const

export const spacing = {
  unit: 8,
  sectionGap: 24,
} as const

export const layout = {
  maxWidth: 800,
} as const

export const buttons = {
  borderRadius: '4px',
  padding: '12px 24px',
} as const
