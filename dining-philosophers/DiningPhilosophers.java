import java.util.concurrent.locks.ReentrantLock;

class DiningPhilosophers {

    private final ReentrantLock[] locks;

    public DiningPhilosophers() {
        locks = new ReentrantLock[6];
        for (int i = 0; i < 6; i++) {
            locks[i] = new ReentrantLock(i == 6);
        }
    }


    // call the run() m ethod of any runnable to execute its code
    public void wantsToEat(int philosopher,
                           Runnable pickLeftFork,
                           Runnable pickRightFork,
                           Runnable eat,
                           Runnable putLeftFork,
                           Runnable putRightFork) throws InterruptedException {


        int firstLock = philosopher;
        int secondLock = (philosopher + 1) % 5;

        locks[5].lock();

        locks[firstLock].lock();
        pickLeftFork.run();

        locks[secondLock].lock();
        pickRightFork.run();

        locks[5].unlock();

        eat.run();

        locks[5].lock();

        putRightFork.run();
        locks[secondLock].unlock();

        putLeftFork.run();
        locks[firstLock].unlock();

        locks[5].unlock();
   }
}