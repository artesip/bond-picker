import { themeQuartz } from 'ag-grid-community';

export const shadcnTheme = themeQuartz.withParams({
  backgroundColor: 'var(--background)',
  foregroundColor: 'var(--foreground)',

  headerBackgroundColor: 'var(--card)',
  headerTextColor      : 'var(--muted-foreground)',

  oddRowBackgroundColor: 'var(--background)',

  borderColor: 'var(--border)',

  rowHoverColor: 'var(--card)',

  selectedRowBackgroundColor: 'var(--accent)',
});