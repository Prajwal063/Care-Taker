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
        .max(32, 'Maximum 32 characters required in description.'),
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
        .max(10, 'Maximum 10 characters required in Phone.'),
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
