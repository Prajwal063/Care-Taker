import express from 'express';
import { getUser, updateUser } from '../controller/user.controller';
import { isAuthenticated } from '../middleware';

const router = express.Router();

router
    .route('/me')
    .get(isAuthenticated, getUser)
    .put(isAuthenticated, updateUser);

export default router;
