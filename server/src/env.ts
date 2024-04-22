import { parseEnv, port, z } from 'znv';

export const {
    PORT,
    MONGO_URL,
    GOOGLE_CLIENT_ID,
    GOOGLE_CLIENT_SECRET,
    CLIENT_URL,
} = parseEnv(process.env, {
    MONGO_URL: {
        schema: z
            .string()
            .url()
            .default('mongodb://localhost:27017/care-taker'),
        description: 'MongoDB database URL',
    },
    PORT: port().default(8000),
    GOOGLE_CLIENT_ID: {
        schema: z.string().min(3),
        description: 'Google OAuth client ID',
    },
    GOOGLE_CLIENT_SECRET: {
        schema: z.string().min(3),
        description: 'Google OAuth client secret',
    },
    CLIENT_URL: {
        schema: z.string().url().default('http://localhost:3000'),
        description: 'Client URL',
    },
});
