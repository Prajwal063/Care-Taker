import type { NextFunction, Request, Response } from 'express';

export { default as errorHandler } from './errorHandler';
export { default as validate } from './validate';

export const isAuthenticated = (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    if (!req.isAuthenticated()) {
        return res.status(401).json({ message: 'Unauthorized' });
    }
    next();
};
