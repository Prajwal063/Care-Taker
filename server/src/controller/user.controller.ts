import type { NextFunction, Request, Response } from 'express';
import AppError from '../helper/AppError';
import { User } from '../models';

export const getUser = async (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    try {
        const { user_id } = req.headers;

        const user = await User.findOne({ _id: user_id });

        res.status(200).json({
            message: 'Fetched user Successfully!',
            data: user,
        });
    } catch (err: any) {
        next(
            new AppError({
                message: err.message || 'Server error occurred!',
                statusCode: err.statusCode || 400,
                stack: err.stack || '',
            })
        );
    }
};

export const updateUser = async (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    try {
        const { id } = req.body;

        const user = await User.findOneAndUpdate(
            { _id: id },
            { ...req.body },
            { new: true }
        );

        res.status(200).json({
            message: 'User updated Successfully!',
            data: user,
        });
    } catch (err: any) {
        next(
            new AppError({
                message: err.message || 'Server error occurred!',
                statusCode: err.statusCode || 400,
                stack: err.stack || '',
            })
        );
    }
};
