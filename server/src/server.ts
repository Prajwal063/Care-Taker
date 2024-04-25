import cookieSession from 'cookie-session';
import cors, { CorsOptions } from 'cors';
import express from 'express';
import helmet from 'helmet';
import morgan from 'morgan';
import passport from 'passport';
import { errorHandler } from './middleware';
import { authRoute, eventRoute, serviceRoute, userRoute } from './routes';

const app = express();
const corsOptions: CorsOptions = {
    origin: '*',
    methods: 'GET,POST,PUT,DELETE',
    credentials: true,
};

app.use(helmet())
    .use(helmet.hidePoweredBy())
    .use(morgan('dev'))
    .use(express.json())
    .use(express.urlencoded({ extended: true }))
    .use(cors(corsOptions))
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
app.use('/user', userRoute);

app.use(errorHandler);

export default app;
