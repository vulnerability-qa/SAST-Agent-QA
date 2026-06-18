import express, { Request, Response } from 'express';

const app = express();

app.get('/greet', (req: Request, res: Response) => {
  const name = req.query.name as string || '';
  // CWE-79: user input reflected directly into HTML without escaping.
  res.send(`<h1>Hello, ${name}!</h1>`);
});

app.listen(3000);
