import { parseEnv, port, z } from 'znv';

export const { PORT, MONGO_URL } = parseEnv(process.env, {
    MONGO_URL: {
        schema: z
            .string()
            .url()
            .default('mongodb://localhost:27017/care-taker'),
        description: 'MongoDB database URL',
    },

    PORT: port().default(8000),
});
