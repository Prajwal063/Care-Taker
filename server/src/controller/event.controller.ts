import type { NextFunction, Request, Response } from 'express';
import AppError from '../helper/AppError';
import { Event } from '../models';

export const getAllEvents = async (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    try {
        const { search } = req.query;

        const events = await Event.find({
            title: { $regex: search || '', $options: 'i' },
        });

        res.status(200).json({
            message: 'Fetched events Successfully!',
            data: events,
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

export const createEvent = async (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    try {
        const event = new Event(req.body);

        await event.save();

        res.status(201).json({
            message: 'Event created Successfully!',
            data: event,
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

export const getEventById = async (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    try {
        const { id } = req.params;

        const event = await Event.findById(id);
        res.status(200).json({
            message: 'Fetched event Successfully!',
            data: event,
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

export const updateEvent = async (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    try {
        const { id } = req.params;

        const event = await Event.findByIdAndUpdate(id, req.body, {
            new: true,
        });

        res.status(200).json({
            message: 'Updated event Successfully!',
            data: event,
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

export const deleteEvent = async (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    try {
        const { id } = req.params;

        await Event.findByIdAndDelete(id);

        res.status(200).json({
            message: 'Deleted event Successfully!',
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

export const registerForEvent = async (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    try {
        const { id } = req.params;
        const { user_id } = req.body;

        const event = await Event.findByIdAndUpdate(id);
        if (!event) {
            throw new Error('Event not found');
        }

        console.log(event.registeredUsers, user_id);
        if (event.registeredUsers.includes(user_id)) {
            throw new Error('User already registered for this event');
        }
        event.registeredUsers.push(user_id);
        await event.save();

        res.status(200).json({
            message: 'Registered for event Successfully!',
            data: event,
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

export const unregisterForEvent = async (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    try {
        const { id } = req.params;
        const { user_id } = req.body;

        const event = await Event.findByIdAndUpdate(id);
        if (!event) {
            throw new Error('Event not found');
        }
        if (!event.registeredUsers.includes(user_id)) {
            throw new Error('User not registered for this event');
        }
        event.registeredUsers = event.registeredUsers.filter(
            (id) => id !== user_id
        );
        await event.save();

        res.status(200).json({
            message: 'Registered for event Successfully!',
            data: event,
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
