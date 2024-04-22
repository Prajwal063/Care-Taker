import { z } from 'zod';

export const CreateServiceType = z.object({
    title: z
        .string({
            required_error: 'Title is required.',
        })
        .min(3, 'Minimum 3 characters required in Title.')
        .max(32, 'Maximum 32 characters required in Title.'),
    description: z
        .string({
            required_error: 'description is required.',
        })
        .min(3, 'Minimum 3 characters required in description.')
        .max(500, 'Maximum 500 characters required in description.'),
    email: z
        .string({
            required_error: 'Email is required.',
        })
        .email('Invalid Email.'),
    phone: z
        .string({
            required_error: 'Phone is required.',
        })
        .min(10, 'Minimum 10 characters required in Phone.')
        .max(10, 'Maximum 10 characters required in Phone.')
        .regex(/^\d+$/, 'Invalid Phone Number.'),
    address: z
        .string({
            required_error: 'Address is required.',
        })
        .min(3, 'Minimum 3 characters required in Address.')
        .max(100, 'Maximum 100 characters required in Address.'),
    picture: z
        .string({
            required_error: 'Picture is required.',
        })
        .url('Invalid Picture URL.'),
});

const EventStatusSchema = z.union([
    z.literal('UPCOMING'),
    z.literal('ONGOING'),
    z.literal('COMPLETED'),
    z.literal('CANCELLED'),
]);

export const CreateEventType = z.object({
    title: z
        .string({
            required_error: 'Title is required.',
        })
        .min(3, 'Minimum 3 characters required in Title.')
        .max(32, 'Maximum 32 characters required in Title.'),
    description: z
        .string({
            required_error: 'description is required.',
        })
        .min(3, 'Minimum 3 characters required in description.')
        .max(500, 'Maximum 500 characters required in description.'),
    email: z
        .string({
            required_error: 'Email is required.',
        })
        .email('Invalid Email.'),
    phone: z
        .string({
            required_error: 'Phone is required.',
        })
        .min(10, 'Minimum 10 characters required in Phone.')
        .max(10, 'Maximum 10 characters required in Phone.')
        .regex(/^\d+$/, 'Invalid Phone Number.'),
    address: z
        .string({
            required_error: 'Address is required.',
        })
        .min(3, 'Minimum 3 characters required in Address.')
        .max(100, 'Maximum 100 characters required in Address.'),
    picture: z
        .string({
            required_error: 'Picture is required.',
        })
        .url('Invalid Picture URL.'),
    date: z
        .string({
            required_error: 'Date is required.',
        })
        .regex(/^\d{2}-\d{2}-\d{4}$/, 'Invalid Date.'),
    time: z
        .string({
            required_error: 'Time is required.',
        })
        .regex(/^\d{2}:\d{2}$/, 'Invalid Time.'),
    price: z
        .number({
            required_error: 'Price is required.',
        })
        .min(0, 'Minimum Price is 0.'),
    status: EventStatusSchema,
});
