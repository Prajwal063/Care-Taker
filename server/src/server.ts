import cookieSession from 'cookie-session';
import cors from 'cors';
import express from 'express';
import helmet from 'helmet';
import morgan from 'morgan';
import passport from 'passport';
import errorHandler from './middleware/errorHandler';
import { authRoute, eventRoute, serviceRoute } from './routes';

const app = express();

app.use(helmet())
    .use(helmet.hidePoweredBy())
    .use(morgan('dev'))
    .use(express.json())
    .use(express.urlencoded({ extended: true }))
    .use(cors())
    .use(
        cookieSession({
            name: 'session',
            keys: ['care-taker-key'],
            maxAge: 24 * 60 * 60 * 100,
        })
    )
    .use(passport.initialize())
    .use(passport.session());

app.get('/', (_, res) => {
    return res.status(200).json({
        message: 'Hello World!',
    });
});

app.use('/auth', authRoute);
app.use('/service', serviceRoute);
app.use('/event', eventRoute);

app.use(errorHandler);

export default app;
