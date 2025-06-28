"use client";
import {
  CssBaseline,
  ThemeProvider,
  createTheme,
  Container,
} from "@mui/material";

const theme = createTheme();

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>
        <ThemeProvider theme={theme}>
          <CssBaseline />
          <Container maxWidth="md" sx={{ mt: 4, mb: 4 }}>
            {children}
          </Container>
        </ThemeProvider>
      </body>
    </html>
  );
}
